#!groovy

node('docker') {

    // This stage is required to provide writing permissions for the Docker instance
    stage('Prepare') {
        def uid = sh returnStdout: true, script: 'id -u'
        def gid = sh returnStdout: true, script: 'id -g'
        sh "echo 'jenkins:x:${uid.trim()}:${gid.trim()}:Jenkins:${env.WORKSPACE}:/bin/bash' > passwd"
    }

    // Setting the passwd file with user information as created in the Prepare step...
    docker.image('golang:1.23.3').inside("-v ${env.WORKSPACE}/passwd:/etc/passwd:ro") {

        stage('Checkout') {
            checkout scm
        }

        stage('Install Module') {
            sh 'go mod tidy'
        }

        // Using a Go tool to format test outputs into a Jenkins-readable XML format.
        stage('Unit Tests') {
            sh 'go install github.com/jstemmer/go-junit-report@v1.0.0'
            sh 'go test -v 2>&1 | go-junit-report > ./testdata/report.xml'
            junit stdioRetention: '', testResults: 'testdata/report.xml'

            if (currentBuild.currentResult == 'UNSTABLE') {
                currentBuild.result = 'FAILURE'
                echo 'The test process unexpectedly turned out to be unstable. Please check test logs.'
            }
        }

        stage("Release") {
            if (env.BRANCH_NAME ==~ /release\/.*/ && currentBuild.currentResult == 'SUCCESS') {
                def version = env.BRANCH_NAME - 'release/'
                def message = "Release version ${version}"
                sh 'git config --replace-all "remote.origin.fetch" "+refs/heads/*:refs/remotes/origin/*"'
                sh 'git fetch origin master'
                sh "git -c user.name='CES Marvin' -c user.email='cesmarvin@cloudogu.com' tag -m '${message}' v${version}"
                sh "git checkout master"
                sh 'git reset --hard origin/master'
                sh "git merge --ff-only v${version}"
                withCredentials([
                    usernamePassword(credentialsId: 'SCM-Manager', usernameVariable: 'AUTH_USR', passwordVariable: 'AUTH_PSW')
                ]) {
                    sh "git -c credential.helper=\"!f() { echo username='\$AUTH_USR'; echo password='\$AUTH_PSW'; }; f\" push --tags origin master"
                    sh "git -c credential.helper=\"!f() { echo username='\$AUTH_USR'; echo password='\$AUTH_PSW'; }; f\" push -d origin ${env.BRANCH_NAME}"
                }
                withCredentials([
                    usernamePassword(credentialsId: 'cesmarvin', usernameVariable: 'AUTH_USR', passwordVariable: 'AUTH_PSW')
                ]) {
                    sh "git -c credential.helper=\"!f() { echo username='\$AUTH_USR'; echo password='\$AUTH_PSW'; }; f\" push --tags https://github.com/scm-manager/goscm"
                }
            }
        }

        stage("Push to GitHub") {
            if (env.BRANCH_NAME == 'master' && currentBuild.currentResult == 'SUCCESS') {
                withCredentials([
                    usernamePassword(credentialsId: 'cesmarvin', usernameVariable: 'AUTH_USR', passwordVariable: 'AUTH_PSW')
                ]) {
                    sh "git -c credential.helper=\"!f() { echo username='\$AUTH_USR'; echo password='\$AUTH_PSW'; }; f\" push https://github.com/scm-manager/goscm HEAD:master"
                }
            }
        }
    }
}