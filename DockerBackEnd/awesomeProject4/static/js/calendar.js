document.addEventListener('DOMContentLoaded', function() {
    var calendarEl = document.getElementById('calendar');
    
    addr = window.location.href.split('/')
    id = addr[addr.length - 1]
    arr = []
    $.ajax({
        type: "GET",
        url: "../api/tasks/"+id,
        success: function (response) {
            if (response != null) {
                for (let i = 0; i < response.length; i++) {
                    arr[i] = response[i];
                }
            }
        }
    })
    .always(function() {
        if (arr.length > 0 ) {
            var calendar = new FullCalendar.Calendar(calendarEl, {
                initialView: 'dayGridMonth',
                events: arr
            });
        } else {
            var calendar = new FullCalendar.Calendar(calendarEl, {
                initialView: 'dayGridMonth'
            });
        }
        calendar.render();
    });
});
