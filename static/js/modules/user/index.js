$(function () {
    InitTable(tableIdIndex, ajaxUrlIndex, paramsSearchIndex, columnsIndex);
});

var ajaxUrlIndex = API + "/user/list";
var ajaxUrlForm = WEB + "/user/form";
var ajaxUrlSave = API + "/user/save";
var ajaxUrlUpdate = API + "/user/update";
var ajaxUrlGetId = "/user/get";
var ajaxUrlDeleteId = "/user/delete";

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
        "data": "Username",
        "title": "Username",
        "sClass": "lcol",
        orderable: false,
        "render": function (data, type, row) {
            return '<a href="javascript:void(0)" data-role="detail" data-id="' + row.Id + '">' + data + '</a>';
        }
    },
    {
        "data": "Name",
        "title": "Name",
        "sClass": "rcol",
        orderable: false
    },
    {
        "data": "Id",
        "title": "Action",
        "sClass": "rcol",
        orderable: false,
        "render": function (data, type, row) {
            return "<a href='javascript:void(0)' class='btn btn-default' onclick=\"EditFormIndex('" + data + "');\"><span title='Ubah' class='glyphicon glyphicon-edit'></span>Edit</a> " +
                " <a href='javascript:void(0)' class='btn btn-danger' onclick=\"DeleteFormIndex('" + row.Id + "');\"><span title='Hapus' class='glyphicon glyphicon-trash'></span>Delete</a>";
        }
    }
];

var paramsSearchIndex = function () {
    var search = $('#searchTextIndex').val();
    var roleId = $("#searchRoleIndex").val();
    return { search: search, role_id: roleId };
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

