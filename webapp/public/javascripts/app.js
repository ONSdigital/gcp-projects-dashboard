$(function() {
  $(".bookmark").on("click", function(event) {
    if ($(event.target).attr("src") == "/svg/bookmark-off.svg") {
      $.post("/addbookmark", { bookmark: event.target.id });
      $(event.target).attr("src", "/svg/bookmark-on.svg?r=" + Math.random());
      $(event.target).attr("alt", "Project bookmarked");
    } else {
      $.post("/removebookmark", { bookmark: event.target.id });
      $(event.target).attr("src", "/svg/bookmark-off.svg?r=" + Math.random());
      $(event.target).attr("alt", "Bookmark project");
    }
  });

  displayDatesInLocalTime();
});

function displayDatesInLocalTime() {
  const dateElements = document.querySelectorAll('.date');

  dateElements.forEach(element => {
    const utcDateStr = element.textContent.trim();
    const localDateStr = parseAndConvertToLocalTime(utcDateStr + ':00');

    element.textContent = localDateStr;
  });
}

function parseAndConvertToLocalTime(utcDateString) {
  const [day, month, year, time] = utcDateString.split(' ');

  const months = {
      Jan: '01', Feb: '02', Mar: '03', Apr: '04', May: '05', Jun: '06',
      Jul: '07', Aug: '08', Sep: '09', Oct: '10', Nov: '11', Dec: '12'
  };

  const formattedUTCDate = `${year}-${months[month]}-${day}T${time}Z`;
  const dateObj = new Date(formattedUTCDate);
  const datePart = dateObj.toLocaleDateString('en-GB', { day: 'numeric', month: 'short', year: 'numeric' });
  const timePart = dateObj.toLocaleTimeString('en-GB', { hour: '2-digit', minute: '2-digit', hour12: false });

  return `${datePart} ${timePart}`;
}
