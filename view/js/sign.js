// this is the id of the form
$("#SignForm").submit(function(e) {

    let url = $(this).attr('action');
    let username = $("#InputUserName").val();
    let password = $("#InputPassword").val();
    let data = new FormData();
    data.append('username', username);
    data.append('password', password);

    $.ajax({
        method: "POST",
        url: url,
        data: data,
        processData: false,
        contentType: false,
        xhrFields: {
            withCredentials: true
        },
        success: function(data){
            HideSignContainer();
        },
        error: function(jqXHR, textStatus, errorThrown){
            $("#AlertWrongParam").show();
        }
    });

    e.preventDefault(); // avoid to execute the actual submit of the form.
});

// click on sign in
$("#SigninBtn").on("click", function(){
    $("#singContainer").show();
});

// click on x
$("#closeSignBtn").on("click", function(){
    HideSignContainer()
});

function HideSignContainer() {
    $("#singContainer").hide();
}