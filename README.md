<div>
  <h1><a href="http://hchat.org">hChat</a></h1>
  <div>
    <h3> The App</h3>
    <p>hChat is a powerful self-hosted tool built with the needs of CUNY students in mind. It is divided hierarchically through all 25 CUNY schools, containing each schools majors, which contain each majors classes. Every school and major contain their own global chatrooms. Within its immense hierarchy hChat encompasses over 86000 different rooms. Due to it's tremendous asynchronicity through Golangs incredibly efficient goroutines, hChat is capable of maintaining thousands of simultaneous connections and conversations between end users, even with its current self-hosted bottleneck. </p>
  </div>
  <hr>
  <div>
    <h3>The Tech Stack</h3>
    <div>
      <h4>React - Frontend</h4>
      <p>The Frontend Functionality and Rendering is written 100% in React. More than anything this choice was made to cope with constant messages and the performance load they could cause. Given React is able to handle live changes highly performatively, this choice was a no-brainer. This was an excellent venture into learning the React Library, updating and managing components, and JSX. </p>
    </div>
    <div>
      <h4>Node/Express - Middleware</h4>
      <p>Node with Express is used to export the web app build over port 3000. This is simply a precautionary measure to prevent unnecessary load in the case the Go server is under high stress.</p>
    </div>
    <div>
      <h4>Go - Backend </h4>
      <p>Go runs the backend simply due to it's concurrent capabilities. Goroutines handle concurrency better than any language on the market, and with a project of this scale, which by nature requires the capability to maintain thousands of potential connections in tens of thousands of rooms, Go was the clear choice. The Go Server is capable of both Websocket connections over the frontend website and potential TCP connections over a desktop client or CLI, while sharing the same pool of messages and users. 
    </div>
    <hr>
  </div>
  <div>
    <h3> The Devs</h3>
    <div>
      <h4>Daniel Volchek</h4>
      <p>Team Lead/Lead Dev</p>
      <p><a href="https://github.com/DanielVolchek">Personal Github Link</a></p>
      <p>I designed and lead development for the functionality of both the GO backend and the React frontend along with organizing the project, team collaboration, and Git Flow. Along the way I extensively learned and researched GO, HTML/CSS/JS/REACT, Git, SQL, and web hosting with NGINX/NODE</p>
    </div>
    <div>
      <h4>Deland Chen</h4>
      <p>Developer</p>
      <p>Researched and assisted design of frontend and backend functionality. Designed and implemented Python scripts to manipulate and update extensive class Database entries. Used Git to maintain code</p>
      <a href="https://github.com/delandchen">Personal Github Link</a>
    </div>
    <div>
    <h4>Anthony Regner</h4>
      <p>Designer/Interface Developer</p>
      <p><a href="https://github.com/A278PlusPi">Personal Github Link</a></p>
      <p>Anthony designed the hChat frontend interface using HTML and CSS, including the proposed layout of the interface. He was willing to learn HTML and CSS for styling websites</p>
    </div>
    <div>
    ...
    </div>
  </div>
</div>
<hr>
<div>
  <h3>On the Docket</h3>
  <p>hChat is still in beta and will be receiving regular updates with bugfixes and new features<p>
  <p>The following is a list of additions/changes ordered by current priority</p>
  <ul>
    <li>HTTPS security</li>
    <li>Bug testing/fixing</li>
    <li>Mobile site</li>
    <li>Active user list</li>
    <li>User accounts through firebase</li>
    <li>Improved logging</li>
    <li>Improved message database querying</li>
    <li>User settings</li>
    <li>Dark theme</li>
    <li>CLI</li>
  </ul>
</div>
</div>
