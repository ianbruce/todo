package persistence

import (
  "github.com/patrickmn/go-cache"

  "github.com/ianbruce/todo/model"
)

type CacheProxy struct {
  Proxied TodoDatabase
  Cache *cache.Cache
}

func (proxy *CacheProxy) GetList(id string) (model.TodoList, error) {
  list, found := proxy.Cache.Get(id)

  if found {
    return list.(model.TodoList), nil
  }

  newList, err := proxy.Proxied.GetList(id)

  if err != nil {
    return newList, err
  }

  proxy.Cache.Set(id, newList, cache.DefaultExpiration)

  return newList, nil
}

func (proxy *CacheProxy) GetLists(query string, limit int) ([]model.TodoList, error) {
  return proxy.Proxied.GetLists(query, limit)
}

func (proxy *CacheProxy) CreateList(list model.TodoList) error {
  return proxy.Proxied.CreateList(list)
}

func (proxy *CacheProxy) CreateTask(listId string, task model.Task) error {
  return proxy.Proxied.CreateTask(listId, task)
}

func (proxy *CacheProxy) UpdateTaskStatus(listId string, taskId string, completion bool) error {
  return proxy.Proxied.UpdateTaskStatus(listId, taskId, completion)
}
