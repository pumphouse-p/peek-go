pipeline {
  agent any

  tools {
    go "golang-1.16.14"
  }

  stages {
    stage('Build') {
      steps {
        echo 'Building...'
        sh "go build"
      }
    }
    stage('Test') {
      steps {
        echo 'Testing...'
        sh "go test ./..."
      }
    }
    stage('Deploy') {
      steps {
        echo 'Deploying...'
      }
    }
  }
}
