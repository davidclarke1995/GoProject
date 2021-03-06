Student ID: G00335563

# GoProject

## Introduction

My name is David Clarke, I am an 3rd year student in Software Development. In one of my modules Graph Theory, I was asked to design a project that is worth 50% of the overall mudule grade. This project contains a Go program that can execute regular expressions against strings over the alphabet 0, 1 and with the special characters ., |, and ∗. We have to create a program that can turn a regular expression into postfix notation, and then into a non deterministic finite automota. It was written in Visual Studio Code in GoLang. 

## User Instructions:

### Downloading the project

The user must download both git and go from the following links if they have not already  https://golang.org/dl/ https://git-scm.com/downloads.
To clone the remote repository to the users local machine, they must enter https://github.com/davidclarke1995/GoProject.git into the command prompt

### Running the program

To run the program in visual studio code, the user must select "Ctrl" and ' to open up the terminal while in the runner class. In the terminal the user must type "go build GoProject.go" to build the program and must type "go run GoProject.go" to compile and run the code.

## About Project

Before starting this project, I read the brief and specs thoroughly and then watched the videos that were posted weekly on the moodle page. Understanding how the project works was essential before starting. The shunt coding video was what I heavily referenced my work on and pulled from. The Thompson algorithm was what the project is based on. 

I began by planning out the workings of the project and deciding how it would work. I am quite familiar with writing in java and decided to lay out my project how I would with a standard java project. This included involving a main class or a "runner". To less complicate things for myself, I had seperate classes run the individual calculations and the functions.

I added more characters (than initially stated in the brief) to use as elements such as "+" and "?", this worked successfully but trying to add other elements let me down. Trying to actually commit my workings was a real issue for me halfway through doing this project as i was getting the error:

"! [rejected] master -> master (non-fast-forward) error: failed to push some refs to 'git@github.com/davidclarke1995/GoProject.git' To prevent you from losing history, non-fast-forward updates were rejected Merge the remote changes (e.g. 'git pull') before pushing again. See the 'Note about fast-forwards' section of 'git push --help' for details."

It suggested that I pull the project down to local repository but if I tried a "git pull" of any kind, I would get the error message:

"Pull is not possible because you have unmerged files. Please, fix them up in the work tree, and then use 'git add/rm ' as appropriate to mark resolution, or use 'git commit -a'."

So I seemed to be going in circles, meanwhile I could continue working on the project without committing but eventually a few days later through StackOverflow, I found a handy feature on git "git reset --hard". My local repository was behind the remote repository but this reset everything and I was able to commit and push again. 

The .gitignore was set up to ignore certain files and directories. I looked up templates such as https://github.com/github/gitignore for the certain elements to have in my .gitignore file and the reasons https://www.quora.com/Whats-the-purpose-of-a-gitignore-file?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa such as .exe as I don't want my executable file to be visible in the github repository.

## Conclusion

Overall I was very happy at the learning process of this project. The videos that my lecturer Dr. Ian McLoughlin had up on moodle were very helpful to understand and there was a lot of content online of Thompsons Algorithm. I found comparing pofix expression to NFA very dificult but solved through help of StackOverflow and some other online material. Also, I had a huge struggle with GitHub not being able to commit successfully but eventually got that issue sorted. I believe the error was because I updated the Readme on GitHub rather than on visual studios and committed so my remote repository and local repository were out of sync and I could not perform a "git pull" request. I think I have learned a lot from this project and look forward to using GoLang again in the future. 

## References

https://web.microsoftstream.com/video/9d83a3f3-bc4f-4bda-95cc-b21c8e67675e
https://stackoverflow.com/questions/35272592/how-are-finite-automata-implemented-in-code?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
https://golang.org/
https://cs.stackexchange.com/questions/30114/converting-nfa-to-regular-expression?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa
https://en.wikipedia.org/wiki/Automata-based_programming
https://cs.stackexchange.com/questions/43845/does-thompsons-algorithm-produce-optimal-nfas
https://stackoverflow.com/questions/20061252/convert-a-nfa-to-regular-expression?utm_medium=organic&utm_source=google_rich_qa&utm_campaign=google_rich_qa