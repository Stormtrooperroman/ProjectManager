$("#send").click(function (e) {
    let name = $("#inputLogin").val();
    let pass = $("#inputPassword").val();
    let send_data = JSON.stringify({
        username: name,
        password: pass
    })
    $.ajax({
        type: "GET",
        url: "../api/login",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            console.log(response);
            if (response['login'] == 'true') {
                // let old = window.location.href;
                window.location.href = "http://localhost:3000/";
            }
        }
    });
    
});