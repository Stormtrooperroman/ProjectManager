$("#update").click(function (e) {

    $('#title').removeClass("is-invalid");
    $('#description').removeClass("is-invalid");


    $('#title').addClass("is-valid");
    $('#description').addClass("is-valid");

    let title = $("#title").val();
    let description = $("#description").val();
    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(title)) {
        is_valid = false
        $('#title').removeClass("is-valid");
        $("#title").addClass("is-invalid")
        // show error
    }
    if (re.test(description)) {
        is_valid = false
        $('#description').removeClass("is-valid");
        $("#description").addClass("is-invalid")
        // show error
    }

    let title_name = $("#title").val();
    let description_val = $("#description").val();
    let color_val = $("#color").val();
    let text_color_val = $("#text_color").val();
    console.log(color_val)
    let send_data = JSON.stringify({
        Name: title_name,
        description: description_val,
        Colour: color_val,
        TextColor: text_color_val
    })

    if(!is_valid){
        return
    }

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