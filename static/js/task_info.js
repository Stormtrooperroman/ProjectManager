let myModal = new bootstrap.Modal(document.getElementById('persons_add'))

$("#persons").click(function (e) {
    myModal.show();
});

let toastElList = [].slice.call(document.querySelectorAll('.toast'))
let toastList = toastElList.map(function (toastEl) {
    return new bootstrap.Toast(toastEl)
})
for(i = 0; i < toastList.length; i++) {
    toastList[i].show()
}

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
    let title_name = $("#title").val();
    let start_date_val = $("#startDate").val();
    let end_date_val = $("#endDate").val();
    let description_val = $("#description").val();
    let person_val = $("#person").val();
    let send_data = JSON.stringify({
        title: title_name,
        startDate: start_date_val,
        endDate: end_date_val,
        description: description_val,
        person: person_val
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
        success: function (response) {
            console.log("Ok")
        }
    });
    
});