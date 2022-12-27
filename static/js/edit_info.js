$("#update").click(function (e) {
    let title_name = $("#title").val();
    let description_val = $("#description").val();
    let color_val = $("#color").val();
    
    console.log(color_val)
    let send_data = JSON.stringify({
        name: title_name,
        description: description_val,
        colour: color_val
    })
    let page_url = window.location.href.split('/')
    let project_id = page_url[page_url.length - 2]
    $.ajax({
        type: "POST",
        url: "../../../api/project/"+project_id,
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            console.log("Ok")
        }
    });
    
});