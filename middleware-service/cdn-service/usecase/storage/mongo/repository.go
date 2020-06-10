package mongo

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/golang/glog"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/gridfs"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"cdn-service/core/repo"
	"cdn-service/env"
)

type repository struct {
	client  *mongo.Client
	database   *mongo.Database
	timeout time.Duration
}

func NewRepository(config env.DataStoreConfig) (repo.ContentRepository, error) {
	repo := &repository{
		timeout: time.Duration(config.Timeout) * time.Second,
	}
	client, err := newMongoClient(config.URL, config.Timeout)
	if err != nil {
		glog.Errorln(err)
		return nil, errors.Wrap(err, "repository.NewRepository")
	}
	repo.client = client
	database := client.Database(config.Database)
	repo.database = database
	return repo, nil
}

func newMongoClient(mongoURL string, mongoTimeout int) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(mongoTimeout)*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURL))
	if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	return client, nil
}


func (r *repository) Download(from string, filename string) ([]byte, error) {
	file, err := ioutil.TempFile("tmp", fmt.Sprintf("%v.%v", from, filename))
	if err != nil {
		glog.Errorln(err)
		return nil, err
	}
	_, err = r.bucketOf(from).DownloadToStreamByName(filename, file); if err != nil {
		glog.Fatalln(err)
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func (r *repository) Upload(to string, filename string, content []byte) (string, error) {
	id, err := r.bucketOf(to).UploadFromStream(filename, bytes.NewBuffer(content)); if err != nil {
		glog.Errorln(err)
		return "", err
	}
	return id.String(), nil
}


func (r *repository) bucketOf(collection string) *gridfs.Bucket {
	bucket, err := gridfs.NewBucket(r.database, &options.BucketOptions{
		Name: &collection,
	}); if err != nil {
		glog.Errorln(err)
		return nil
	}
	return bucket
}