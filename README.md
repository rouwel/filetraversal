Hello There. 

Have you ever tried to find a file or folder  that was so well hidden on your C drive or root directory that you just gave up, well then, this is the repo for you, by the time we(if you decide to help) are done this program should be able to genrate the path to any given folder or file as long as it exists.

For now, this repository holds a skeleton of a program that should be able to sift through multiple folders  in a drive and locate a specific folder or file.

Right now all it does is identifies folders in the same directory it is in.
To see for yourself 
<pre>
clone https://github.com/rouwel/filetraversal.git
</pre>

Then proceed to add a folder of your choosing into the same folder you have cloned.
After this run the command
<pre>
go run . "foldername"
</pre>
Output shoud be 
<pre>
Directory Found
</pre>
If the folder does not exist the output should be.
<pre>
Path does not exist:GetFileAttributesEx front: The system cannot find the file specified.
</pre>