$(document).ready(() => {
    $('#subdomains').DataTable({
        "pagingType": "simple",
        "language": {
            "lengthMenu": "Amount: _MENU_",
            "info": "_PAGE_ of _PAGES_",
            "infoEmpty": "0 of 0",
            "infoFiltered": ""
        }
    });
    $('.dataTables_length').addClass('bs-select');
});