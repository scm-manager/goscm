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
    }
}