let myModal = new bootstrap.Modal(document.getElementById('persons_add'))

$("#persons").click(function (e) {
    myModal.show();
});

let toastElList = [].slice.call(document.querySelectorAll('.user_toast'))
let toastList = toastElList.map(function (toastEl) {
    return new bootstrap.Toast(toastEl)
})
for(i = 0; i < toastList.length; i++) {
    toastList[i].show()
}

$("#add").click(function (e) {
    let new_person = $("#new_person").val()
    let new_person_id = new_person.split(" ")[2]
    let person_val = []
    let persons_data = [].slice.call(document.querySelectorAll('.toast-body'))
    persons_data.forEach(toaster => {
        person_val.push(toaster["innerText"].split(" ")[2])
    });
    if (new_person != "Выберетите сотрудника" && !(person_val.includes(new_person_id))) {
        test = 
        `<div class="toast align-items-center user_toast user_toast"  role="alert" aria-live="assertive" aria-atomic="true" data-bs-autohide='false'>
            <div class="d-flex">
            <div class="toast-body">
            `+new_person+`
            </div>
            <button type="button" class="btn-close me-2 m-auto delete" data-bs-dismiss="toast" aria-label="Close"></button>
            </div>
        </div>`

        $("#all_persons").append(test);
        let toastElList = [].slice.call(document.querySelectorAll('.user_toast'))
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


    $('#title').removeClass("is-invalid");
    $('#startDate').removeClass("is-invalid");
    $('#endDate').removeClass("is-invalid");


    $('#title').addClass("is-valid");
    $('#startDate').addClass("is-valid");
    $('#endDate').addClass("is-valid");


    let title_name = $("#title").val();
    let start_date_val = $("#startDate").val();
    let end_date_val = $("#endDate").val();
    let description_val = $("#description").val();
    let start_date = new Date(start_date_val).getTime()
    let end_date = new Date(end_date_val).getTime()
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


    if (start_date > end_date) {
        is_valid = false
        $('#startDate').removeClass("is-valid");
        $("#startDate").addClass("is-invalid");
        $('#endDate').removeClass("is-valid");
        $("#endDate").addClass("is-invalid");
        // show error
    }

    console.log(end_date+" "+ start_date)

    if(!is_valid){
        return
    }




    let person_val = []
    let persons_data = [].slice.call(document.querySelectorAll('.toast-body'))
    persons_data.forEach(toaster => {
        person_val.push(toaster["innerText"].split(" ")[2])
    });
    let is_fin = false

    if ($('#is_finished').is(":checked")){
        is_fin = true
    }
    let send_data = JSON.stringify({
        title: title_name,
        start: start_date_val,
        end: end_date_val,
        text: description_val,
        person_mas: person_val,
        is_finished: is_fin
    })
    let page_url = window.location.href.split('/')
    let task_id = page_url[page_url.length - 1]
    let project_id = page_url[page_url.length - 3]




    $.ajax({
        type: "POST",
        url: "../../../api/project/"+project_id+"/task/"+task_id,
        data: send_data,
        contentType: "application/json",
        dataType: "json",
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done_update'))
                toast.show()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });
    
});

$("#delete").click(function (e) {

    let page_url = window.location.href.split('/')
    let task_id = page_url[page_url.length - 1]
    let project_id = page_url[page_url.length - 3]
    $.ajax({
        type: "POST",
        url: "../../../api/from_project/"+project_id+"/del_task/"+task_id,
        statusCode:{
            200:function() {
                let toast = new bootstrap.Toast(document.getElementById('done_del'))
                toast.show()

                location.reload()
            },
            500: function() {
                let toast = new bootstrap.Toast(document.getElementById('fail'))
                toast.show()
            }
        }
    });

});