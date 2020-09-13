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
            SetUserCookie(data.user)
            window.location.reload();
        },
        error: function(jqXHR, textStatus, errorThrown){
            console.log(jqXHR)
            console.log(textStatus)
            console.log(errorThrown)
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

function SetUserCookie(user) {
    let sub = [["text", "subOid"], ["text", "subOuniquename"],["text", "subOnickname"],["text", "subOdescription"]];
    split_strs(user, sub)
    document.cookie = "OwnerList=" + JSON.stringify(user) + ";Path=/;domain=.dcreater.com;Max-Age=2592000;SameSite=Lax";
}