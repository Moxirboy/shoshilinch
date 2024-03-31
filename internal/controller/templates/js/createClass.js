


function CreateClass(event){
    event.preventDefault();
    const name =document.getElementById('name').value;
    const password=document.getElementById('password').value
    const body =JSON.stringify({
        name:name,
        password:password,
    })
    console.log(body)
    fetch('/v1/class/create',{
        method:"POST",
        headers:{
            'Content-Type':'application/json'
        },
        body:body
    })
    .then(response=>{
        if (!response.ok){
            throw new Error('Network response was not ok');
        }
        return response.json();
    })
        .then(data=>{
            console.log(data)
            document.getElementById("myDropdown").reset();

        })
    .catch(error=>{
        console.error('error:',error)
    })
}
function goBack(event) {
    event.preventDefault();
    window.location.href = '/teacher';
}