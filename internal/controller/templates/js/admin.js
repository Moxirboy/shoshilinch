

function SubmitEvent(event){
    event.preventDefault(); // Prevent default form submission

    const first_name = document.getElementById('first-name').value;
    const last_name = document.getElementById('last-name').value;
    const phone_number = document.getElementById('phone-number').value;
    const password = document.getElementById('password').value;
    const body = JSON.stringify({
        first_name: first_name,
        last_name: last_name,
        phone_number: phone_number,
        password: password
    });
    console.log(body)

    fetch('/v1/admin/create_teacher',{
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: body
    })
    .then(response=>{
        if(!response.ok){
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data=>{
        console.log(data);
    })
    .catch(error=>{
        console.error('Error:', error);
    })
}
function GetTeachers(){
    fetch('/v1/admin/teachers',{
        method: 'GET',
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response=>{
        if(!response.ok){
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
    .then(data=>{
        console.log(data);
        const table = document.getElementById('teachers-table');
        
        data.forEach(teacher=>{
            const row = table.insertRow(-1);
            const cell1 = row.insertCell(0);
            const cell2 = row.insertCell(1);
            const cell3 = row.insertCell(2);
            const cell4 = row.insertCell(3);
            const cell5 = row.insertCell(4);
            cell1.innerHTML = teacher.id;
            cell2.innerHTML = teacher.first_name;
            cell3.innerHTML = teacher.last_name;
            cell4.innerHTML = teacher.phone_number;
            cell5.innerHTML = teacher.password;
        })
    })
    .catch(error=>{
        console.error('Error:', error);
    })
}

// Call the function when the page loads
window.onload = function() {
    GetTeachers();
};