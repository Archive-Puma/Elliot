// -----------------------------------------------------------------------------
// PACKAGE
// -----------------------------------------------------------------------------

package scanner

// -----------------------------------------------------------------------------
// IMPORTS
// -----------------------------------------------------------------------------

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/table"

	"github.com/cosasdepuma/elliot/pkg/app"
	"github.com/cosasdepuma/elliot/pkg/logger"
	"github.com/cosasdepuma/elliot/pkg/modules"
)

// -----------------------------------------------------------------------------
// SUBCOMMAND DEFINITION
// -----------------------------------------------------------------------------

// Nmap TODO
var Nmap = modules.Module{
	Flag: args,
	Run:  Run,
}

// -----------------------------------------------------------------------------
// ARGUMENTS
// -----------------------------------------------------------------------------

var (
	args    = flag.NewFlagSet("nmap", flag.ContinueOnError)
	host    = args.String("host", "127.0.0.1", "Host a escanear")
	port    = args.String("port", "1-1000", "Puertos a escanear")
	timeout = args.Duration("timeout", 1*time.Second, "Timeout")
	threads = args.Int("threads", 1000, "Number of threads")
	retries = args.Int("retries", 1, "Number of retries if a port is close")
)

// -----------------------------------------------------------------------------
// PUBLIC METHODS
// -----------------------------------------------------------------------------

// Run TODO
func Run(core *app.Core) {
	// Get the arguments
	flag.Parse()
	// Parse "port" argument
	ports, err := portSplitter(*port)
	if err != nil {
		log.Fatalf("[!] %s\n", err)
		return
	}
	// Resolve host
	ipaddr, err := net.ResolveIPAddr("", *host)
	if err != nil {
		log.Fatalf("[!] Bad host: %s", *host)
		return
	}
	ip := ipaddr.String()
	if *host != ip {
		*host = fmt.Sprintf("%s (%s)", *host, ip)
	}
	// Scan ports
	logger.Command("nmap", "Scanning %d port(s) from %s", len(ports), *host)
	maxThreads := len(ports) * *retries
	if maxThreads < *threads {
		logger.Warning("Optimizing the number of threads from %d to %d", *threads, maxThreads)
		*threads = maxThreads
	}
	result := pingTCPRange(ip, ports, *timeout, *threads, *retries)

	// Filter open ports
	openPorts := make([]int, 0)
	for port, open := range result {
		if open {
			openPorts = append(openPorts, port)
		}
	}

	// ! Beautify ports (just dev)

	t := table.NewWriter()
	t.SetStyle(table.StyleRounded)
	if len(openPorts) == 0 {
		t.AppendHeader(table.Row{"No open ports"})
	} else {
		sort.Ints(openPorts)
		t.AppendHeader(table.Row{"protocol", "port", "status"})
		for _, port := range openPorts {
			t.AppendRow(table.Row{"tcp", port, "open"})
		}
	}
	logger.Put("\n%s", t.Render())
}

// -----------------------------------------------------------------------------
// PRIVATE METHODS
// -----------------------------------------------------------------------------

// * Utilities

func portSplitter(raw string) ([]int, error) {
	portMap := map[int]struct{}{} // force ports to be unique

	// split ports
	blocks := strings.Split(raw, ",") // split blocks
	for _, block := range blocks {
		rang := strings.SplitN(block, "-", 2) // split ranges
		low, err := strconv.Atoi(rang[0])     // lower port
		if err != nil {
			return nil, errors.New("Bad port specification")
		}
		// append ports
		if len(rang) == 1 {
			portMap[low] = struct{}{} // single port
		} else {
			high, err := strconv.Atoi(rang[1]) // higher port
			if err != nil {
				return nil, errors.New("Bad port specification")
			}
			for port := low; port <= high; port++ {
				portMap[port] = struct{}{} // multiple ports
			}
		}
	}
	// retrieve ports
	ports := make([]int, 0, len(portMap))
	for port := range portMap {
		ports = append(ports, port)
	}
	return ports, nil
}

// * Workflow

func pingTCP(host string, port int, timeout time.Duration, retries int) bool {
	try := 0                                 // current try
	status := false                          // open port flag
	addr := fmt.Sprintf("%s:%d", host, port) // format address
	for try < retries && !status {
		conn, err := net.DialTimeout("tcp", addr, timeout) // connect to given addres with a timeout
		if err == nil {
			conn.Close()  // close the open connection
			status = true // port is open
		} else {
			logger.Put(err.Error())
		}
		try++
	}
	return status
}

func pingTCPRange(host string, ports []int, timeout time.Duration, threads int, retries int) map[int]bool {
	results := map[int]bool{}
	// concurrency: wait group
	wg := sync.WaitGroup{}
	wg.Add(len(ports))
	// concurrency: locks
	mlock := sync.Mutex{}
	tlock := make(chan struct{}, threads)
	// concurrent ping

	for _, port := range ports {
		tlock <- struct{}{} // lock a thread
		// scan a port
		go func(port int) {
			defer wg.Done()
			status := pingTCP(host, port, timeout, retries)
			// concurrency: mutual exclusion
			mlock.Lock()
			// progress
			// logger.Progress("Checking %d", port)
			results[port] = status
			mlock.Unlock()
			<-tlock // unlock a thread
		}(port)
	}
	wg.Wait() // wait for all goroutines to finish

	return results
}
