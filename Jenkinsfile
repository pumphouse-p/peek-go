pipeline {
  agent any

  tools {
    go "golang-1.16.14"
  }

  stages {
    stage('Build') {
      steps {
        echo 'Building...'
        sh "go build ./cmd/peek-go/"
      }
    }
    stage('Test') {
      steps {
        echo 'Testing...'
        sh "go test ./..."
      }
    }
    stage('Build Release') {
      steps {
        echo 'Building Image...'
        sh "buildah bud -t quay.io/deparris/peek-go:jenkins"
      }
    }
    stage('Publish Release') {
      steps {
        echo 'Publishing Image...'
      }
    }
  }
}
