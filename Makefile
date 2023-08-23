# Test load-test with wrk
loadtest:
	wrk -c50 -d10s http://localhost:8085/load-test --latency

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