const dropDown = document.getElementById('myDropdown');

// Function to fetch and populate dropdown options
function getAllClasses() {
    fetch('/v1/class/get_all', {
        method: "GET",
        headers: {
            'Content-Type': 'application/json'
        },
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data => {
        data.data.forEach(item => {
            const option = document.createElement('option');
            option.value = item.name;
            option.text = item.name;
            dropDown.appendChild(option);
        });
    })
    .catch(error => {
        console.log('Error:', error);
    });
}

// Function to create an exam
function createExam(event) {
    event.preventDefault();
    const name = document.getElementById('myDropdown').value;
    const body = JSON.stringify({ name: name });

    fetch('/v1/exam/create', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: body
    })
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not right');
        }
        return response.json();
    })
    .then(data => {
        const successMessage = document.getElementById('successMessage');
        successMessage.textContent = 'Exam created successfully';
        successMessage.classList.add('show'); // Show the success message
        setTimeout(() => {
            successMessage.classList.remove('show'); // Hide the success message after 3 seconds
        }, 3000);
        document.getElementById('myDropdown').value = ''; // Clear dropdown selection
    })
    .catch(error => {
        console.error('Error', error);
    });
}

// Function to navigate back
function goBack() {
    window.location.href = '/teacher';
}

// Call the function to populate dropdown options on page load
getAllClasses();
