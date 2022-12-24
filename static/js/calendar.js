document.addEventListener('DOMContentLoaded', function() {
    var calendarEl = document.getElementById('calendar');
    
    addr = window.location.href.split('/')
    id = addr[addr.length - 1]
    arr = []
    $.ajax({
        type: "GET",
        url: "../api/tasks/"+id,
        success: function (response) {
            for (let i = 0; i < response.length; i++) {
                arr[i] = response[i];
                console.log(arr[i])
            }
        },
        
    })
    .always(function() {
        var calendar = new FullCalendar.Calendar(calendarEl, {
            initialView: 'dayGridMonth',
                events: arr
        });
        calendar.render();
    });
});
