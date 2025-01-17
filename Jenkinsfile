pipeline {
    agent { label 'centos' }

    environment {
        OUTPUT_DIR = 'build'
        BRANCH_NAME = ''
        GITHUB_REPO = "class-scheduling-system/table-install-cli"
        GITHUB = credentials('github-token')
        TAG_VERSION = ''
    }

    stages {
        stage('获取分支信息') {
            steps {
                script {
                    // 从 git 获取当前分支名称
                    BRANCH_NAME = sh(script: "git rev-parse --abbrev-ref HEAD", returnStdout: true).trim()
                    echo "当前分支是 ${BRANCH_NAME}"
                }
            }
        }

        stage('签出代码') {
            when {
                expression { BRANCH_NAME == 'master' || BRANCH_NAME == 'HEAD' }
            }
            steps {
                echo "拉取代码中..."
                checkout scm
            }
        }

        stage('项目打包') {
            when {
                expression { BRANCH_NAME == 'master' || BRANCH_NAME == 'HEAD' }
            }
            steps {
                script {
                    // 生成时间戳版本号
                    TAG_VERSION = sh(script: "date +%Y%m%d%H%M%S", returnStdout: true).trim()
                    echo "清空构建目录..."
                    sh "rm -rf ${WORKSPACE}/build/${TAG_VERSION}"
                    sh "mkdir -p ${WORKSPACE}/build/${TAG_VERSION}"

                    echo "编译项目..."
                    def platforms = [
                        [os: 'windows', arch: 'amd64'],
                        [os: 'windows', arch: 'arm64'],
                        [os: 'linux', arch: 'amd64'],
                        [os: 'linux', arch: 'arm64'],
                        [os: 'darwin', arch: 'amd64'],
                        [os: 'darwin', arch: 'arm64']
                    ]

                    platforms.each { platform ->
                        def outputFile = "${WORKSPACE}/build/${TAG_VERSION}/monitor-dashboard-${platform.os}-${platform.arch}${platform.os == 'windows' ? '.exe' : ''}"
                        sh """
                            GOOS=${platform.os} GOARCH=${platform.arch} go build -o ${outputFile} ./main.go
                            echo "完成打包: ${outputFile}"
                        """
                    }
                }
            }
        }

        stage('收集构建产物') {
            when {
                expression { BRANCH_NAME == 'master' || BRANCH_NAME == 'HEAD' }
            }
            steps {
                script {
                    echo "版本号: ${TAG_VERSION}"
                    echo "收集所有构建产物..."
                    sh "ls -al ${WORKSPACE}/build/${TAG_VERSION}"
                }
            }
        }

        stage('发布到 GitHub Release') {
            when {
                expression { BRANCH_NAME == 'master' || BRANCH_NAME == 'HEAD' }
            }
            steps {
                script {
                    def RELEASE_NAME = "RELEASE-${TAG_VERSION}"
                    def CHANGELOG_FILE = "CHANGELOG.md"

                    echo "生成 CHANGELOG..."
                    sh """
                        echo "# DESCRIPTION" > ${CHANGELOG_FILE}
                        echo "\$(git log -1 --pretty=format:'%B') \n" >> ${CHANGELOG_FILE}
                        echo "# CHANGELOG" >> ${CHANGELOG_FILE}
                        git log --pretty=format:'- %s (@%an) [%h]' >> ${CHANGELOG_FILE}
                        echo "" >> ${CHANGELOG_FILE}
                    """

                    def loginStatus = sh(
                        script: "echo $GITHUB | gh auth login --with-token && gh auth status | grep 'Logged in'",
                        returnStatus: true
                    )

                    if (loginStatus != 0) {
                        error "❌ GitHub CLI 登录失败，请检查 GITHUB_TOKEN 是否正确！"
                    } else {
                        echo "✅ GitHub CLI 登录成功"
                    }

                    sh """
                        echo "发布到 GitHub Release..."
                        gh release create ${TAG_VERSION} ${OUTPUT_DIR}/${TAG_VERSION}/** \\
                            --repo ${GITHUB_REPO} \\
                            --title "${RELEASE_NAME}" \\
                            --notes-file ${CHANGELOG_FILE}
                    """
                }
            }
        }
    }

    post {
        success {
            echo "流水线执行成功 🎉"
        }
        failure {
            echo "流水线执行失败，请检查日志 ❌"
        }
    }
}