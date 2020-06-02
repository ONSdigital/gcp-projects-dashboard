$(function() {
  $(".bookmark").on( "click", function(event) {
    $.post("/bookmark", { bookmark: event.target.id });
  });
});
