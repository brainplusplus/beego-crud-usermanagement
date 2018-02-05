var NOTIF_TIMEOUT = 2000;
var NOTIF_TOP_CENTER = "top-center";
var NOTIF_TOP_LEFT = "top-left";
var NOTIF_TOP_RIGHT = "top-right";
var NOTIF_BOTTOM_CENTER = "bottom-center";
var NOTIF_BOTTOM_LEFT = "bottom-left";
var NOTIF_BOTTOM_RIGHT = "bottom-right";
var NOTIF_DEFAULT = "primary";
var NOTIF_DANGER = "danger";
var NOTIF_WARNING = "warning";
var NOTIF_SUCCESS = "success";

function GetValueOrDefault(value, defaultValue) {
    return typeof value == "undefined" ? defaultValue : value;
}

function GeneralNotif(message, status, timeout, pos) {
    message = GetValueOrDefault(message, "");
    status = GetValueOrDefault(status, NOTIF_DEFAULT);
    timeout = GetValueOrDefault(timeout, NOTIF_TIMEOUT);
    pos = GetValueOrDefault(pos, NOTIF_TOP_CENTER);

    UIkit.notify({
        message: message,
        status: status,
        timeout: timeout,
        pos: pos
    });
}

function CustomNotif(message, options = { status: NOTIF_DEFAULT, timeout: NOTIF_TIMEOUT, pos: NOTIF_TOP_CENTER }) {
    GeneralNotif(message, options.status, options.timeout, options.pos);
}

function DangerNotif(message, options = { timeout: NOTIF_TIMEOUT, pos: NOTIF_TOP_CENTER }) {
    GeneralNotif(message, NOTIF_DANGER, options.timeout, options.pos);
}

function WarningNotif(message, options = { timeout: NOTIF_TIMEOUT, pos: NOTIF_TOP_CENTER }) {
    GeneralNotif(message, NOTIF_WARNING, options.timeout, options.pos);
}
function SuccessNotif(message, options = { timeout: NOTIF_TIMEOUT, pos: NOTIF_TOP_CENTER }) {
    GeneralNotif(message, NOTIF_SUCCESS, options.timeout, options.pos);
}

function DefaultNotif(message, options = { timeout: NOTIF_TIMEOUT, pos: NOTIF_TOP_CENTER }) {
    GeneralNotif(message, NOTIF_DEFAULT, options.timeout, options.pos);
}

function ajaxGet(e, p, cb) {
    axios.get(e, {
        params: p
    })
        .then(cb)
        .catch(function (error) {
            console.log(error);
            DangerNotif(error);
        })
}

function ajaxPost(e, p, cb) {
    axios.post(e, p)
        .then(cb)
        .catch(function (error) {
            console.log(error);
            DangerNotif(error);
        })
}

//api
function ajaxGetApi(e, p, cb) {
    ajaxGet(API + e, p, cb);
}

function ajaxPostApi(e, p, cb) {
    ajaxPost(API + e, p, cb);
}

function ajaxGetApiFinance(e, p, cb) {
    ajaxGet(API_FINANCE + e, p, cb);
}

function ajaxPostApiFinance(e, p, cb) {
    ajaxPost(API_FINANCE + e, p, cb);
}

function ajaxGetApiHris(e, p, cb) {
    ajaxGet(API_HRIS + e, p, cb);
}

function ajaxPostApiHris(e, p, cb) {
    ajaxPost(API_HRIS + e, p, cb);
}


function ajaxGetApiGeneral(e, p, cb) {
    ajaxGet(API_GENERAL + e, p, cb);
}

function ajaxPostApiGeneral(e, p, cb) {
    ajaxPost(API_GENERAL + e, p, cb);
}

function ajaxGetApiPM(e, p, cb) {
    ajaxGet(API_PM + e, p, cb);
}

function ajaxPostApiPM(e, p, cb) {
    ajaxPost(API_PM + e, p, cb);
}

//web
function ajaxGetWeb(e, p, cb) {
    ajaxGet(WEB + e, p, cb);
}

function ajaxPostWeb(e, p, cb) {
    ajaxPost(WEB + e, p, cb);
}

function ajaxGetWebFinance(e, p, cb) {
    ajaxGet(WEB_FINANCE + e, p, cb);
}

function ajaxPostWebFinance(e, p, cb) {
    ajaxPost(WEB_FINANCE + e, p, cb);
}

function ajaxGetWebHris(e, p, cb) {
    ajaxGet(WEB_HRIS + e, p, cb);
}

function ajaxPostWebHris(e, p, cb) {
    ajaxPost(WEB_HRIS + e, p, cb);
}


function ajaxGetWebGeneral(e, p, cb) {
    ajaxGet(WEB_GENERAL + e, p, cb);
}

function ajaxPostWebGeneral(e, p, cb) {
    ajaxPost(WEB_GENERAL + e, p, cb);
}

function ajaxGetWebPM(e, p, cb) {
    ajaxGet(WEB_PM + e, p, cb);
}

function ajaxPostWebPM(e, p, cb) {
    ajaxPost(WEB_PM + e, p, cb);
}


var MODAL_BLOCKUI;
var MODAL_FOR_ALL;
var MODAL_FOR_ALL_ID = "modal_for_all";

showModalContent(title, content){
    MODAL_FOR_ALL = UIkit.modal(MODAL_FOR_ALL_ID);
    $("#" + MODAL_FOR_ALL_ID + " .uk-modal-title").html(title);
    $("#" + MODAL_FOR_ALL_ID + " .uk-modal-content").html(content);
    MODAL_FOR_ALL.show();
}

hideModalContent(){
    if (typeof MODAL_FOR_ALL != "undefined")
    {
        MODAL_FOR_ALL.hide();
    }
}

showModalBlockUI(content){
    MODAL_BLOCKUI = UIkit.modal.blockUI(content);
}

hideModalBlockUI(){
    if (typeof MODAL_BLOCKUI != "undefined") {
        MODAL_BLOCKUI.hide();
    }
}

showModalAjaxGet(title, url, p, cb){
    axios.get(url, {
        params: p
    })
        .then(function (response) {
            showModalContent(title, response);
            if (typeof cb != "undefined") {
                cb();
            }
        })
        .catch(function (error) {
            console.log(error);
            DangerNotif(error);
        })
}

showModalAjaxPost(title, url, p, cb){
    axios.post(url,p)
        .then(function (response) {
            showModalContent(title, response);
            if (typeof cb != "undefined") {
                cb();
            }
        })
        .catch(function (error) {
            console.log(error);
            DangerNotif(error);
        })
}

function checkLogin() {
    /*$.get("${urlApp}/checkSession",function(resultSession){
     resultSession = parseJson(resultSession);
     if(!resultSession.isLogin){
     alert("Session login anda sudah habis ! silahkan login lagi");
     window.location.href="${urlApp}/login";
     }
     });*/
}
function checkActivity(timeout, interval, elapsed) {
    if ($.active) {
        elapsed = 0;
        $.active = false;
        $.get('checkEBOnlineSession');
    }
    if (elapsed < timeout) {
        elapsed += interval;
        setTimeout(function () {
            checkActivity(timeout, interval, elapsed);
        }, interval);
    } else {
        window.location.href = 'login'; // Redirect to login page
    }
}

function parseJson(data) {
    if (typeof data == "string")
        data = jQuery.parseJSON(data);

    return data;
}
function FormLoadByAjaxGet(formid,url, p, cb) {
    ajaxGet(url, p, function (data) {
        FormLoadByDataUsingName(formid, data, cb);
    });
}

function FormLoadByAjaxPost(formid, url, p, cb) {
    ajaxPost(url, p, function (data) {
        FormLoadByDataUsingName(formid, data, cb);
    });
}

function FormLoadByData(data, formid, attrKey, cb) {
    attrKey = typeof attrKey == "undefined" ? "name" : attrKey;
    $.each(data, function (key, value) {
        //console.log(key+' '+ value);
        if ($('#' + formid + ' [' + attrKey + '="' + key + '"]').attr('data-role') == "datepicker") {
            //window.console.log('1');
            $('#' + formid + ' [' + attrKey + '="' + key + '"]').val(value);
        }
        else if ($('#' + formid + ' [' + attrKey + '="' + key + '"]').attr('data-role') == "combobox") {
            //window.console.log('2');
            $('#' + formid + ' [' + attrKey + '="' + key + '"]').val(value);
        }
        else if ($('#' + formid + ' [' + attrKey + '="' + key + '"]').attr('data-role') == "numeric") {
            //window.console.log('3');
            $('#' + formid + ' [' + attrKey + '="' + key + '"]').val(value);
        }
        else if ($('#' + formid + ' [' + attrKey + '="' + key + '"]').attr('data-role') == "checkbox") {
            //window.console.log('3');
            if (value == "1" || value == 1 || value == "true" || value == true)
                $('#' + formid + ' [' + attrKey + '="' + key + '"]').prop('checked', true);
            else
                $('#' + formid + ' [' + attrKey + '="' + key + '"]').removeAttr('checked');
        }
        else {
            //window.console.log('4');
            $('#' + formid + ' [' + attrKey + '="' + key + '"]').val(value);
        }
    });
    if (typeof cb != "undefined") {
        cb();
    }
}
function FormLoadByDataUsingID(data, formid, cb) {
    FormLoadByData(data, formid, "id", cb);
}
function FormLoadByDataUsingName(data, formid, cb) {
    FormLoadByData(data, formid, "name", cb);
}