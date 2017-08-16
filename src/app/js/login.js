$(function(){
    //centerLoginBoxVertically();
    //centerLoginBoxHorizontally();
});

function centerLoginBoxVertically(){
    var loginWrapper = $("#login-wrapper");
    var divHeight = parseInt(loginWrapper.outerHeight());
    var marginTop = parseInt((divHeight / 2) * -1);

    loginWrapper.css("margin-top", marginTop+"px");
}

function centerLoginBoxHorizontally(){
    var loginWrapper = $("#login-wrapper");
    var divWidth = parseInt(loginWrapper.outerWidth());
    var marginLeft = parseInt((divWidth / 2) * -1);

    loginWrapper.css("margin-left", marginLeft+"px");
}