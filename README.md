# hChat
<div>
  <h2> Hunter Codefest 2021 - Class chat room application </h2>
  <p> With this application we aim to create a simple TCP chatroom to send messages back and forth between clients.<br>
    The distinguishing feature of our chatroom will be its groups. We will have groups and subgroups of majors as
    demonstrated in the following list: </p>
</div>
<div>
  <p> hChat hierarchy: </p>
  <ul>
    <li>
      Schools (Hunter for now, then CUNY, then other schools)
      <ul>
        <li>School chat room</li>
        <li>
          Majors
          <ul>
            <li>Major chat room</li>
            <li>Class chat rooms</li>
          </ul>
        </li>
    </li>
  </ul>
</div>
<div>
  <hr>
  <h3> Tasks:<br><sub>(edit x into box to check off)</sub> </h3>
  <ul>
    <li>
      Client side: 
      <ul>
        <li> - [ ] Build UI
          <ul>
            <li> - [ ] Create groups in sidebar</li>
            <li> - [ ] Create chat room display</li>
            <li> - [ ] Create and send input from text bar</li>
          </ul>
        </li>
        <li> - [ ] Enter and store username</li>
        <li> - [ ] Send client enter/exit chatroom</li>
        <li> - [ ] Validate text input (client side)</li>
        <li> - [ ] Send text input to server</li>
        <li> - [ ] Receive text input</li>
        <li>- [ ] Display server messages to user</li>
      </ul>
    </li>
    <li>Server side: 
      <ul>
        <li> - [ ] Connect clients to server</li>
        <li> - [ ] Temporarily store client IP</li>
        <li> - [ ] Seperate clients into chatrooms</li>
        <li> - [ ] Validate text input (server side)</li>
        <li> - [ ] Process text</li>
        <li> - [ ] Return processed text to every IP in indicated chatroom</li>
      </ul>
    </li>
  </ul>
</div>
<div>
  <hr>
  <h3>Text processing: </h3>
  <p>Text processing can be done by encoding the messages as follows: 
    <br>
    #00000 - To indicate room
    <br>
    _username - To indicate user
    <br>
    :message - To indicate message
    <br>
    The full syntax should look something like this:
    <br>
    #00000_username:message
    <br>
    #00000_username with no message can indicate entering and exiting chatroom
  </p>
</div>
