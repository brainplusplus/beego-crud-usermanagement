$(function () {
    InitTable(tableIdIndex, ajaxUrlIndex, paramsSearchIndex, columnsIndex);
});

var ajaxUrlIndex = API + "/role/list";
var ajaxUrlForm = WEB + "/role/form";
var ajaxUrlSave = API + "/role/save";
var ajaxUrlUpdate = API + "/role/update";
var ajaxUrlGetId = "/role/get";
var ajaxUrlDeleteId = "/role/delete";

var tableIdIndex = "tableIndex";
var formIdIndex = "formIndex";
var searchBoxIdIndex = "searchBoxIndex";

var columnsIndex = [
    {
        "data": "Id",
        "title": "ID",
        "sClass": "ecol x20",
        orderable: false
    },
    {
        "data": "Name",
        "title": "Name",
        "sClass": "lcol",
        orderable: false,
        "render": function (data, type, row) {
            return '<a href="javascript:void(0)" data-role="detail" data-id="' + row.Id + '">' + data + '</a>';
        }
    },
    {
        "data": "Notes",
        "title": "Notes",
        "sClass": "rcol",
        orderable: false
    },
    {
        "data": "Id",
        "title": "Action",
        "sClass": "rcol",
        orderable: false,
        "render": function (data, type, row) {
            return "<a href='javascript:void(0)' class='btn btn-default' onclick=\"EditFormIndex('" + data + "');\"><span title='Ubah' class='glyphicon glyphicon-edit'></span>Edit</a> ";
 //              + " <a href='javascript:void(0)' class='btn btn-danger' onclick=\"DeleteFormIndex('" + row.Id + "');\"><span title='Hapus' class='glyphicon glyphicon-trash'></span>Delete</a>";
        }
    }
];

var paramsSearchIndex = function () {
    var search = $('#searchTextIndex').val();
    return { search: search};
}




function ShowSearchTableIndex() {
    showCustomModal(searchBoxIdIndex);
}
function SearchTableIndex() {
    ReloadGridIndex();
    hideAllModal();
}

function ReloadGridIndex() {
    ReloadGrid(tableIdIndex);
}

