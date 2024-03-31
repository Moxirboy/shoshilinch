


function SubmitEvent(event) {
    event.preventDefault(); // Prevent default form submission
  
    const phone_number = document.getElementById('phone_number').value;
    const password = document.getElementById('password').value;
    console.log(phone_number, password);
    const body = JSON.stringify({
      phone_number: phone_number,
      password: password
    });
    console.log(body)
    fetch('/v1/auth/login', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: body
    })
    .then(response => {
      console.log(response)
      if (!response.ok) {
        throw new Error('Network response was not ok');
      } 
      
      return response.json();
    })
    .then(data => {
      // Save token to localStorage
    console.log(data);
      localStorage.setItem('token', data.token);
  
      // Check role
      if(data.role === 'teacher') {
        // Redirect to /dash
        window.location.href = '/teacher';
      } else if (data.role === 'student') {
        // Redirect to /dashboard
        window.location.href = '/student';
      }else if (data.role === 'admin') {
        // Redirect to /dashboard
        window.location.href = '/admin';
      } else {
        
        console.log('Role is not teacher or student');
      }
    })
    .catch(error => {
      console.error('Error:', error);
    });
   
  }
  

