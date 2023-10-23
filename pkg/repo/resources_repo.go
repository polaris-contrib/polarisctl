/**
 * Tencent is pleased to support the open source community by making Polaris available.
 *
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 *
 * Licensed under the BSD 3-Clause License (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://opensource.org/licenses/BSD-3-Clause
 *
 * Unless required by applicable law or agreed to in writing, software distributed
 * under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
 * CONDITIONS OF ANY KIND, either express or implied. See the License for the
 * specific language governing permissions and limitations under the License.
 */
package repo

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/0226zy/polarisctl/pkg/entity"
)

type RepoOption func(repo *ResourceRepo)

func WithWriter(writer entity.Writer) RepoOption {
	return func(repo *ResourceRepo) {
		repo.writer = writer
	}
}
func WithParser(parser *entity.ResponseParse) RepoOption {
	return func(repo *ResourceRepo) {
		repo.parser = parser
	}

}
func WithFile(fileName string) RepoOption {
	return func(repo *ResourceRepo) {
		repo.rsFile = fileName
	}
}

func WithParam(value string) RepoOption {
	return func(repo *ResourceRepo) {
		repo.client.queryParam = value
	}
}

func WithMethod(method string) RepoOption {
	return func(repo *ResourceRepo) {
		repo.method = method
	}
}

// ResourceWrite 创建/删除/修改 资源
type ResourceRepo struct {
	client *ApiClient
	method string
	rsFile string
	writer entity.Writer
	parser *entity.ResponseParse
}

// NewResourceRepo 查询操作
func NewResourceRepo(url string, options ...RepoOption) *ResourceRepo {
	ret := &ResourceRepo{
		client: NewApiClient(url),
		writer: entity.NewTableWriter(),
	}
	for _, option := range options {
		option(ret)
	}
	return ret
}

// Method set http method:GET/PUT/POST/PUT/DEL
func (rsRepo *ResourceRepo) Method(value string) *ResourceRepo {
	rsRepo.method = value
	return rsRepo
}

// Build execute
func (rsRepo ResourceRepo) Build() {
	if rsRepo.method == "GET" {
		rsRepo.get()
		return
	}
	rsRepo.write()
}

// write put/post/del resources
func (rsRepo ResourceRepo) write() {
	jsonFile, err := os.Open(rsRepo.rsFile)
	if err != nil {
		fmt.Printf("[polarisctl err] open rs files failed:%v\n", err)
		os.Exit(1)
	}

	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		os.Exit(1)
	}

	body := []byte{}
	httpCode := 0

	if rsRepo.method == "POST" {
		httpCode, body = rsRepo.client.Post(bytes.NewBuffer(jsonData))
	} else if rsRepo.method == "PUT" {
		httpCode, body = rsRepo.client.Put(bytes.NewBuffer(jsonData))
	} else if rsRepo.method == "DELETE" {
		httpCode, body = rsRepo.client.Delete(bytes.NewBuffer(jsonData))
	} else {
		fmt.Printf("[polarisctl internal sys err] unkown method:%s\n", rsRepo.method)
		os.Exit(1)
	}

	if 200 != httpCode {
		rsRepo.writer.Write(&entity.HttpFailed{strconv.Itoa(httpCode), string(body)})
		return
	}
	response := rsRepo.parser.Parse(body)
	rsRepo.writer.Write(response)
}

// get query resources
func (rsRepo ResourceRepo) get() {
	httpCode, body := rsRepo.client.Get()
	if 200 != httpCode {
		rsRepo.writer.Write(&entity.HttpFailed{strconv.Itoa(httpCode), string(body)})
		return
	}
	response := rsRepo.parser.Parse(body)
	rsRepo.writer.Write(response)
}
