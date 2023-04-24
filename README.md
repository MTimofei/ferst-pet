<h1>Ferst Pet Projecr</h1> 

<h2>Microservice for registration, identification and authorization.</h2>

<h2>Ready :</h2>
<p>
<h3>Registration</h3> 
Password hashing using the argon2 algorithm
<p>
<h3>Communication</h3>
Transfer of public keys from the server to the client-server, using the gRPC API
</p>
<p>
<h3>Authentication</h3>
Authorization by a pair of login password, and a refresh token 
The refresh token is signed using the ES256 algorithm
</p>
<h2>In the plans for the future :</2>
<p>
<h3>Authorization</h3>
Using the OAuth 2.0 protocol with an asynchronous encoding algorithm.
</p>
<h3>Communication</h3>
Server synchronization<br>
Implement sending requests for keys at time intervals<br>
</p>
<h3> Safety</h3>
Generating new keys at time intervals for each encryption algorithm
</p>