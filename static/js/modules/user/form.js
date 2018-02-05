function InitFormCreateIndex() {
    var optionsForm = GetOptionsForm(requestFormCreateIndex, responseFormCreateIndex);
    InitForm(formIdIndex, optionsForm);
}

function InitFormUpdateIndex() {
    var optionsForm = GetOptionsForm(requestFormUpdateIndex, responseFormUpdateIndex);
    InitForm(formIdIndex, optionsForm);
}



function AddFormIndex() {
    console.log(ajaxUrlForm);
    showModalAjaxGet("Add", ajaxUrlForm, {}, function (result) {
        setFormAction(formIdIndex, ajaxUrlSave);
        InitFormCreateIndex();
    });
}

function EditFormIndex(id) {
    ajaxGetApi(ajaxUrlGetId + "?id=" + id, {}, function (response) {
        response = parseJson(response);
        //response = response.data;
        showModalAjaxGet("Edit", ajaxUrlForm + "?id="+id, {}, function (result) {
            setFormAction(formIdIndex, ajaxUrlUpdate);
            InitFormUpdateIndex();
            FormLoadByDataUsingName(response.Data, formIdIndex);
            console.log("TEST");
            setTimeout(function () {
                $("#"+formIdIndex+" .formRoles").each(function ()
                {
                    var roleId = $(this).val();
                    $.each(response.Data.Roles, function( idxRole, valRole ) {
                        if(roleId==valRole.Id){
                            console.log("MASUK");
                            $("#formRoles-"+roleId).prop('checked', true);
                        }
                    });
                });
            },0);

        });
    });
}

function DeleteFormIndex(id) {
    var text = "Do you want to delete this data?";
    confirmDialog(text, function () {
        ajaxGetApi(ajaxUrlDeleteId + "?id=" + id, {}, function (response) {
            response = parseJson(response);
            if (response.Success) {
                SuccessNotif("Delete data is successfully !");
                ReloadGridIndex();
                hideAllModal();
            } else {
                DangerNotif(response.Message);
            }
        });
    });
}

function submitFormIndex(){
       var form = $("#"+formIdIndex);
       form.parsley().validate();
       if (form.parsley().isValid()){
            form.submit();
       }
}


var responseFormCreateIndex = function (response, statusText, xhr, $form) {
    response = parseJson(response);
    if (response.Success) {
        SuccessNotif("Save data is successfully !");
        ReloadGridIndex();
        hideAllModal();
    } else {
        DangerNotif(response.Message);
    }
};

var requestFormCreateIndex = function (formData, jqForm, options) {
    return true;
}

var responseFormUpdateIndex = function (response, statusText, xhr, $form) {
    console.log(response);
    response = parseJson(response);
    if (response.Success) {
        SuccessNotif("Update data is successfully !");
        ReloadGridIndex();
        hideAllModal();
    } else {
        DangerNotif(response.Message);
    }
};

var requestFormUpdateIndex = function (formData, jqForm, options) {
    return true;
}
var secondIdModal = "secondModal";
function ShowModalSecond() {
    showCustomModal(secondIdModal, {modal:false,style:""});
}