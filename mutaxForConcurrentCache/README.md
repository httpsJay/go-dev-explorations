# Concurrent Cache in Go

A simple, thread-safe cache implementation in Go that supports setting and getting values with optional expiration durations.

## Data Structures

### `cacheValue`

Holds the actual data you want to store in the cache.

```
|-----------------|
| value           | -> The actual value (any type) you want to store in the cache.
| expiration      | -> The time when the value will expire.
|-----------------|
```

### `ConcurrentCache`

Main structure for the cache that uses mutex for concurrent read/write access.

```
|-----------------|
| mu              | -> Mutex for concurrent read/write access.
| cache           | -> Map from string (key) to cacheValue.
|-----------------|
```

## Functions and Methods

- `NewConcurrentCache()`: Initializes and returns a new instance of `ConcurrentCache` with an empty map.
- `Set(key, value, expiration)`: Adds or updates a value in the cache with a given expiration time.
- `Get(key)`: Retrieves a value from the cache using the given key. If the key doesn't exist or the value has expired, it returns `nil`.

## Flow of Execution

1. **Initialization**: Create a new cache instance.

   ```go
   main() -> NewConcurrentCache() -> ConcurrentCache{}
   ```

2. **Setting Data**: Set data in the cache with expiration durations.

   ```go
   main() -> cache.Set("key1", "data1", 5*time.Second)
   main() -> cache.Set("key2", "data2", 10*time.Second)
   ```

3. **Data Retrieval**: Retrieve and print the data.

   ```go
   main() -> cache.Get("key1") -> ("data1", true)
   main() -> cache.Get("key2") -> ("data2", true)
   ```

4. **Waiting for Expiration**: Wait for 6 seconds to let the data with `key1` expire.

   ```go
   main() -> time.Sleep(6 * time.Second)
   ```

5. **Retrieving Expired Data**: Try to retrieve the expired data (should not be found).

   ```go
   main() -> cache.Get("key1") -> (nil, false)
   ```

## Important Notes

- The `Set` method uses the write lock because it modifies the cache.
- The `Get` method also uses the write lock, especially when deleting an expired value. This ensures atomicity and avoids potential race conditions.

---