import service from '@/utils/request'

// @Summary 新增爬虫任务
// @Router /spider/task [post]
export const createSpiderTask = (data) => {
  return service({
    url: '/spider/task',
    method: 'post',
    data: data,
  })
}

// @Summary 分页获取爬虫任务列表
// @Router /spider/taskList [post]
export const getSpiderTaskList = (data) => {
  return service({
    url: '/spider/taskList',
    method: 'post',
    data: data,
  })
}

// @Summary 删除爬虫任务
// @Router /spider/task [delete]
export const deleteSpiderTask = (data) => {
  return service({
    url: '/spider/task',
    method: 'delete',
    data: data,
  })
}

// @Summary 获取爬虫任务信息
// @Router /spider/taskById [post]
export const getSpiderTaskInfo = () => {
  return service({
    url: '/spider/taskById',
    method: 'post',
  })
}

// @Summary 更新爬虫任务
// @Router /spider/taskStatus [put]
export const updateSpiderTaskStatus = (data) => {
  return service({
    url: '/spider/task',
    method: 'put',
    data: data,
  })
}

// @Tags system
// @Summary 更新爬虫相关配置
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body sysModel.System true
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /spider/updateConfig [post]
export const updateSpiderConfig = (data) => {
  return service({
    url: '/spider/updateConfig',
    method: 'post',
    data,
  })
}

// @Tags systrm
// @Summary 获取爬虫配置文件内容
// @Security ApiKeyAuth
// @Produce  application/json
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /spider/getConfig [post]
export const getSpiderConfig = () => {
  return service({
    url: '/spider/getConfig',
    method: 'post',
  })
}

// @Summary 启动爬虫任务
// @Router /spider/start [post]
export const startSpiderTask = (data) => {
  return service({
    url: '/spider/start',
    method: 'post',
    data,
  })
}

// @Summary 根据字典类型获取字典项
// @Router /spider/findDictByType [post]
export const findDictByType = (data) => {
  return service({
    url: '/spider/findDictByType',
    method: 'post',
    data,
  })
}

// @Summary 创建爬虫字典项
// @Router /spider/dictionary [post]
export const createSpiderDictionary = (data) => {
  return service({
    url: '/spider/dictionary',
    method: 'post',
    data,
  })
}

// @Summary 更新爬虫字典项
// @Router /spider/dictionary [put]
export const updateSpiderDictionary = (data) => {
  return service({
    url: '/spider/dictionary',
    method: 'put',
    data,
  })
}

// @Summary 删除爬虫字典项
// @Router /spider/dictionary [delete]
export const deleteSpiderDictionary = (data) => {
  return service({
    url: '/spider/dictionary',
    method: 'delete',
    data,
  })
}

// @Summary 获取爬虫字典项
// @Router /spider/dictionary [get]
export const getSpiderDictionary = (params) => {
  return service({
    url: '/spider/dictionary',
    method: 'get',
    params
  })
}
