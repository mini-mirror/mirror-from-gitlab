# This project is used to experiement how to mirror from gitlab to github

## gitlab project:

https://<GITLAB-URL>/mirror-gitlab

## github project:

https://github.com/mini-mirror/mirror-gitlab.git

## mirror gitlab project to github

```bash
#clone the gitlab project
git clone git@<GITLAB-URL>/mirror-gitlab.git

#mirror to github
cd mirror-gitlab
git push --mirror https://github.com/mini-mirror/mirror-gitlab.git

#change the remote repo to github
git remote set-url origin https://github.com/mini-mirror/mirror-gitlab.git