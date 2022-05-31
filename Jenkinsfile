pipeline {
    agent {
        docker {
            image 'golang:latest' 
        }
    }
    stages {
        stage('Get') { 
            steps {
                sh 'go get -u' 
            }
        }
    }
}