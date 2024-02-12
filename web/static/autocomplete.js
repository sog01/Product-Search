const inp = document.getElementById("searchInput")
let currentFocus
inp.addEventListener("input", function (e) {
  currentFocus = -1
})
inp.addEventListener("keydown", function (e) {
  var x = document.getElementById(this.id + "autocomplete-list")
  if (x) x = x.getElementsByTagName("div")
  if (e.keyCode == 40) {
    /*If the arrow DOWN key is pressed,
        increase the currentFocus variable:*/
    currentFocus++
    /*and and make the current item more visible:*/
    addActive(x)
  } else if (e.keyCode == 38) {
    //up
    /*If the arrow UP key is pressed,
        decrease the currentFocus variable:*/
    currentFocus--
    /*and and make the current item more visible:*/
    addActive(x)
  } else if (e.keyCode == 13 && currentFocus > -1) {
    /*If the ENTER key is pressed, prevent the form from being submitted,*/
    e.preventDefault()
    if (currentFocus > -1) {
      /*and simulate a click on the "active" item:*/
      if (x) x[currentFocus].click()
    }
  }
  function addActive(x) {
    /*a function to classify an item as "active":*/
    if (!x) return false
    /*start by removing the "active" class on all items:*/
    removeActive(x)
    if (currentFocus >= x.length) currentFocus = 0
    if (currentFocus < 0) currentFocus = x.length - 1
    /*add class "autocomplete-active":*/
    x[currentFocus].classList.add("autocomplete-active")
  }
  function removeActive(x) {
    /*a function to remove the "active" class from all autocomplete items:*/
    for (var i = 0; i < x.length; i++) {
      x[i].classList.remove("autocomplete-active")
    }
  }
})
