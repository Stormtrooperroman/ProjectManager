let myModal = new bootstrap.Modal(document.getElementById('persons_add'))

$("#persons").click(function (e) {
    myModal.show();
});


$("#add").click(function (e) {
    let new_person = $("#new_person").val()
    if (new_person != "Выберетите сотрудника") {
        console.log(new_person);
        test = 
        `<div class="toast align-items-center"  role="alert" aria-live="assertive" aria-atomic="true" data-bs-autohide='false'>
            <div class="d-flex">
            <div class="toast-body">
            `+new_person+`
            </div>
            <button type="button" class="btn-close me-2 m-auto delete" data-bs-dismiss="toast" aria-label="Close"></button>
            </div>
        </div>`

        $("#all_persons").append(test);
        let toastElList = [].slice.call(document.querySelectorAll('.toast'))
        let toastList = toastElList.map(function (toastEl) {
            return new bootstrap.Toast(toastEl)
        })

        for(i = 0; i < toastList.length; i++) {
            toastList[i].show()
        }

    }
});

$(document).on("click", ".delete", function(){
    $($(this).parent().parent()).remove();
  });

$("#update").click(function (e) {


    $('#taskName').removeClass("is-invalid");
    $('#startDate').removeClass("is-invalid");
    $('#endDate').removeClass("is-invalid");
    $('#description').removeClass("is-invalid");


    $('#taskName').addClass("is-valid");
    $('#startDate').addClass("is-valid");
    $('#endDate').addClass("is-valid");
    $('#description').addClass("is-valid");

    let title_name = $("#taskName").val();
    let start_date_val = $("#startDate").val();
    let end_date_val = $("#endDate").val();
    let description_val = $("#description").val();

    const re = new  RegExp('(^\\s+$|^$)')
    is_valid = true;
    if (re.test(title_name)) {
        is_valid = false
        $('#title').removeClass("is-valid");
        $("#title").addClass("is-invalid")
        // show error
    }

    if (re.test(start_date_val)) {
        is_valid = false
        $('#startDate').removeClass("is-valid");
        $("#startDate").addClass("is-invalid")
        // show error
    }
    if (re.test(end_date_val)) {
        is_valid = false
        $('#endDate').removeClass("is-valid");
        $("#endDate").addClass("is-invalid")
        // show error
    }

    if (re.test(description_val)) {
        is_valid = false
        $('#description').removeClass("is-valid");
        $("#description").addClass("is-invalid")
        // show error
    }

    if(!is_valid){
        return
    }

    let person_val = []
    let persons_data = [].slice.call(document.querySelectorAll('.toast-body'))
    console.log(persons_data)
    persons_data.forEach(toaster => {
        person_val.push(toaster["innerText"].split(" ")[2])
    });

    console.log(person_val)
    let send_data = JSON.stringify({
        title: title_name,
        start: start_date_val,
        end: end_date_val,
        text: description_val,
        person_mas: person_val
    })
    let page_url = window.location.href.split('/')
    let project_id = page_url[page_url.length - 3]
    $.ajax({
        type: "POST",
        url: "../../../api/project/"+project_id+"/add_task/",
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        success: function (response) {
            console.log("Ok")
        }
    });
    
});