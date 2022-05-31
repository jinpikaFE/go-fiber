pipeline {
    agent {
        docker {
            image 'golang:latest' 
        }
    }
    stages {
        stage('SetEnv') { 
            steps {
                sh 'go env -w GOPROXY=https://goproxy.cn,direct'
                sh 'go env -w GO111MODULE=on'
                sh 'go env -w GOARCH=amd64'
            }
        }
        stage('Get') { 
            steps {
                sh 'go get -u' 
            }
        }
    }
}