<h1>Ferst Pet Projecr</h1> 

<h2>Microservice for registration, identification and authorization.</h2>
<p>
<h3>Steak of Technology</h3> 
Golandg<br>
REST API<br>
gRPC<br>
JWT<br>
Argon2<br>
Rsa<br>
ECDA<br>
mySQL<br>
<p>


<h2>Implemented:</h2>
<p>
<h3>Registration</h3> 
Password hashing using the argon2 algorithm<br>
<p>
<h3>Communication</h3>
Transfer of public keys from the server to the client-server, using the gRPC API<br>
Implement sending requests for keys at time intervals<br>
</p>
<p>
<h3>Authentication</h3>
Authorization by a pair of login password, and a refresh token <br>
The refresh token is signed using the ES256 algorithm
</p>
<h3> Safety</h3>
Generating new keys at time intervals for each encryption algorithm<br>
</p>
