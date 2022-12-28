$("#send").click(function (e) {
    let name = $("#inputLogin").val();
    let pass = $("#inputPassword").val();
    let send_data = JSON.stringify({
        login: name,
        pass: pass
    })
    $.ajax({
        type: "POST",
        url: "../api/login",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            console.log(response);
            console.log(response['login']);
            if (response['login'] == "true") {
                // let old = window.location.href;
                console.log(window.location);
                window.location = "http://localhost:3001";
                console.log(response['login']);
            }
        }
    });
    
});