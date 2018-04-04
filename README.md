# GoProject

## Introduction

My name is David Clarke, I am an 3rd year student in Software Development. In one of my modules Graph Theory, I was asked to design a project that is worth 50% of the overall mudule grade. This project contains a Go program that can execute regular expressions against strings over the alphabet 0, 1 and with the special characters ., |, and ∗. We have to create a program that can turn a regular expression into postfix notation, and then into a non deterministic finite automota. It was written in Visual Studio Code in GoLang. 

## User Instructions

To run the program in visual studio code, the user must select "Ctrl" and ' to open up the terminal. In the terminal the user must type "go build GoProject.go" to build the program and must type "go run GoProject.go" to compile and run the code.

## Project Workings

Before starting this project, I read the brief and specs thoroughly and then watched the videos that were posted weekly on the moodle page. Understanding how the project works was essential before starting. The shunt coding video was what I heavily referenced my work on and pulled from. The thompson algorithm was what the project is based on. I added another character to use as an element, "+", this worked successfully but trying to add other elements let me down. Trying to actually commit my workings was a real issue for me halfway through doing this project as i was getting the error:

<<<<<<< HEAD
! [rejected] master -> master (non-fast-forward) error: failed to push some refs to 'git@github.com/davidclarke1995/GoProject.git' To prevent you from losing history, non-fast-forward updates were rejected Merge the remote changes (e.g. 'git pull') before pushing again. See the 'Note about fast-forwards' section of 'git push --help' for details.

It suggested that I pull the project down to local repository but if I tried a "git pull" of any kind, I would get the error message:

Pull is not possible because you have unmerged files. Please, fix them up in the work tree, and then use 'git add/rm ' as appropriate to mark resolution, or use 'git commit -a'.

So I seemed to be going in circles, meanwhile I could continue working on the project without committing but eventually a few days later through StackOverflow, I found a handy feature on git "git reset --hard". My local repository was behind the remote repository but this reset everything and I was able to commit and push again. 
=======
! [rejected]        master -> master (non-fast-forward)
error: failed to push some refs to 'git@github.com/davidclarke1995/GoProject.git'
To prevent you from losing history, non-fast-forward updates were rejected
Merge the remote changes (e.g. 'git pull') before pushing again.  See the
'Note about fast-forwards' section of 'git push --help' for details.

It suggested that I pull the project down to local repository but if I tried a "git pull" of any kind, I would get the error message:

Pull is not possible because you have unmerged files.
Please, fix them up in the work tree, and then use 'git add/rm <file>'
as appropriate to mark resolution, or use 'git commit -a'.

So I seemed to be going in circles, meanwhile I could continue working on the project without committing but eventually a few days later through StackOverflow, I found a handy feature on git "git reset --hard". My local repository was behind the remote repository but this reset everything and I was able to commit and push again.
  

>>>>>>> 110af1e3c1945c7d7d157870a35b88a4b4a4a6b9
