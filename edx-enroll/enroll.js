var system = require('system')
var page = require('webpage').create()

var username = system.args[1]
var password = system.args[2]

page.onConsoleMessage = function(msg) {
  system.stderr.writeLine(msg)
  system.stderr.flush()
}

page.onNavigationRequested = function(url) {
  console.log('Navigating to ' + url)
}

// Invoked after Enroll button is clicked
function courseEnrolled(status) {
  if (status == 'success') {
    console.log('Course Enrolment Successful')
  } else {
    console.log('Course Enrolment Failure')
  }
  phantom.exit()
}


// Invoked after Dashboard is loaded.
function dashboardLoaded(status) {
  console.log(status, page.url)
  if (status == 'success') {
    console.log('Login successful')
    page.onLoadFinished = null
    console.log('Loading course page')
    page.open('https://www.edx.org/course/introduction-programming-java-part-1-uc3mx-it-1-1x-0', function(status) {
      if (status == 'success') {
        console.log('course page loaded')
        page.evaluate(function() {
          console.log('enrolling for course')
          document.querySelector('input.email-opt-in').checked = false
          document.querySelector('div.course-enroll-action a.js-enroll-btn').click()
        })
      } else {
        console.log('course load unsuccessful')
      }
    })
  }
}

// Login to Coursera.
page.open('https://courses.edx.org/login', function(status) {
  if (status !== 'success') {
    console.log('Unable to access network');
  } else {
    page.onLoadFinished = dashboardLoaded
    page.evaluate(function(username, password) {
      console.log('Logging in…');

      document.querySelector("#login-email").value = username
      document.querySelector("#login-password").value = password
      document.querySelector(".login-button").click();

      console.log("Login submitted!")
    }, username, password)
  }
})

