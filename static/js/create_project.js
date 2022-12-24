$("#update").click(function (e) {
    let title_name = $("#title").val();
    let description_val = $("#description").val();
    let color_val = $("#color").val();
    
    console.log(color_val)
    let send_data = JSON.stringify({
        title: title_name,
        description: description_val,
        color: color_val
    })
    $.ajax({
        type: "POST",
        url: "../../../api/new_project/",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            console.log("Ok")
        }
    });
    
});