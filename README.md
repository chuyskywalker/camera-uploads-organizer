# "Camera Uploads" Organizer

Got a "Camera Uploads" folder from Dropbox? Organize it into sane folders!

Dropbox allows for simple uploading of photos, video, etc from your devices. It's a wonderful feature to get all your media into one location. However, a single folder with _thousands_ of files in it is impossible to actually use. This application will move all your media into folders organized by year and month.

Download the [latest release](https://github.com/chuyskywalker/camera-uploads-organize/releases), drop it in the "Camera Uploads" folder, and run the application. Ta da, organized into `yyyy/mm/img` directory structure.

## F.A.Q.

### What are some other options?
- [Daniel Andrade's Pythong Organizer](http://www.danielandrade.net/2014/02/16/script-to-organize-dropboxs-camera-upload-folder/)
- [EXIFmover](http://www.brianklug.org/2012/12/fixing-dropboxs-camera-upload-organization-mess-exifmover/) _Kinda sorta; more geared towards device organization than dates)_
- [organize-dropbox-camera-uploads](https://github.com/nickhammond/organize-dropbox-camera-uploads)

### Why not those other options?
Two reasons:

First, the other options had some odd behaviors I didn't want or simply didn't function the way I desired. Daniel's script put things in folders with month names instead of dates, which doesn't strike my fancy. EXIFmover doesn't really do what I want. Finally, `organize-dropbox-camera-uploads`, required some level of technicality (dropbox api integration) that I thought was too overbearing.

Secondly, these other options almost all required that the end-user install a programming environment on their machines to run the tool (python, ruby, etc). For the ease of use, I wanted something that people could just drop in the right place, run, and be done with it.

### Can I use this, like, anywhere?
Actually, yeah. Well, mostly. If you run this in a different folder, any files that match the naming pattern `yyyy-mm-dd...` will be caught and organized up. So, sure.
