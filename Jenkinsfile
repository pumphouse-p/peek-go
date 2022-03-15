pipeline {
  agent any

  tools {
    go "golang-1.16.14"
  }

  parameters {
    string(name: 'IMAGE_REGISTRY_USERNAME', defaultValue: '', description: 'Username for target image registry')
    password(name: 'IMAGE_REGISTRY_PASSWORD', defaultValue: '', description: 'Password for target image registry')
  }

  stages {
    state('Debug') {
      steps {
        echo "IMAGE_REGISTRY_USERNAME = ${params.IMAGE_REGISTRY_USERNAME}"
        echo "IMAGE_REGISTRY_PASSWORD = ${params.IMAGE_REGISTRY_PASSWORD}"
      }
    }
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
        sh "buildah push --creds ${params.IMAGE_REGISTRY_USERNAME}:${params.IMAGE_REGISTRY_PASSWORD} quay.io/deparris/peek-go:jenkins"
      }
    }
  }
}
