function CreateTest(event){
    event.preventDefault()
    document.getElementById('question').disabled = true;
}

function extend(){
    var newcontent =document.createElement('div');

    newcontent.innerHTML='<input type="text" name="question" placeholder="Question"><input type="text" name="answer" placeholder="Answer">';
 
    var deleteButton = document.createElement('button');
  deleteButton.textContent = 'Delete';
  deleteButton.style.marginLeft = '10px';
  newcontent.appendChild(deleteButton);

    var ContentContainer=document.getElementById('question');
    ContentContainer.appendChild(newcontent);

    deleteButton.addEventListener('click', function() {
        ContentContainer.removeChild(newcontent);
      });
}
document.getElementById('extendButton').addEventListener('click', extend);
