loadtest:
	wrk -c50 -d10s http://localhost:8085/load-test --latency

demo-1-db:
	wrk -c50 -d10s http://localhost:8085/latest-members-db --latency

demo-1-redis:
	wrk -c50 -d10s http://localhost:8085/latest-members-redis --latency

demo-1-redis-v2:
	wrk -c50 -d10s http://localhost:8085/latest-members-redis-v2 --latency