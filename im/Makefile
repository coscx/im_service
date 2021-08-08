all:im

im: im_server.go im.go subscriber.go connection.go client.go peer_client.go group_client.go room_client.go customer_client.go route.go app_route.go protocol.go message.go group_manager.go group.go set.go config.go monitoring.go sio.go  storage_rpc.go channel.go storage_message.go route_message.go user.go rpc.go device.go group_message_deliver.go relationship.go relationship_pool.go
	go build -ldflags "-X main.VERSION=2.0.0 -X 'main.BUILD_TIME=`date`' -X 'main.GO_VERSION=`go version`' -X 'main.GIT_COMMIT_ID=`git log --pretty=format:"%h" -1`' -X 'main.GIT_BRANCH=`git rev-parse --abbrev-ref HEAD`'" -o im im_server.go im.go subscriber.go connection.go client.go peer_client.go group_client.go room_client.go customer_client.go route.go app_route.go protocol.go message.go group_manager.go group.go set.go config.go monitoring.go sio.go storage_rpc.go channel.go storage_message.go route_message.go user.go rpc.go device.go group_message_deliver.go relationship.go relationship_pool.go





main.test:im_test.go user.go relationship.go relationship_pool.go protocol.go message.go config.go
	go test -c  im_test.go user.go relationship.go relationship_pool.go protocol.go message.go config.go


clean:
	rm -f im
