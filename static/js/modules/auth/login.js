
$(document).ready(function () {
    WarningNotif("Login first");

    $('#login_form').on('submit', function (e) {
        e.preventDefault()

        ajaxPostApi('/auth/login', {
            Username: $('#Username').val(),
            Password: $('#Password').val()
        }, function (response) {
            response = parseJson(response);
            //response = response.data
            console.log(response);
            if (response.Success == true) {
                localStorage.setItem("TOKEN", response.JwtCode);
                console.log(localStorage.getItem("TOKEN"))
                var apiToken = response.Data;
                SuccessNotif(response.Message, { pos: NOTIF_BOTTOM_CENTER })

                setTimeout(function () {
                    ajaxGetWeb('/auth/savesession?token=' + apiToken, {},function (response) {
                        console.log(response);
                        window.location = WEB + '/user/index'
                    });
                }, 1000)
            } else {
                DangerNotif(response.Message, { pos: NOTIF_BOTTOM_CENTER });
            }
        })

    })
})