# Test load-test with wrk 10ms
loadtest-10ms:
	wrk -c50 -d10s http://localhost:8085/load-test-10ms --latency

# Test load-test with wrk 300ms
loadtest-500ms:
	wrk -c50 -d10s http://localhost:8085/load-test-300ms --latency

# Query from Database
demo-1-db:
	wrk -c50 -d10s http://localhost:8085/latest-members-db --latency

# Caching with Redis
demo-1-redis:
	wrk -c50 -d10s http://localhost:8085/latest-members-redis --latency

# Optimized with MGET, MSET
demo-2-redis:
	wrk -c50 -d10s http://localhost:8085/latest-members-redis-v2 --latency

# Optimized with MGET, MSET and Memory
demo-2-mem-redis:
	wrk -c50 -d10s http://localhost:8085/latest-members-redis-mem --latency