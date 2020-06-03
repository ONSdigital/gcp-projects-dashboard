$(function() {
  $(".bookmark").on( "click", function(event) {
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
});
