package main

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"sync"

	"github.com/sirupsen/logrus"
	"github.com/yandex-cloud/go-genproto/yandex/cloud/ai/stt/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

var (
	// token = `t1.9euelZqYz4rJmo3Ll5SXy8-Ui5GLx-3rnpWaz8rIjJKXk42Pl82Uz8eQk5bl8_clOGJt-e8KJh9q_d3z92VmX2357womH2r9.0j0PUaj4B3ix3fIDQbY4onObsdxXwk1o2i5cGhIgeH7KCIxw-yTdGR7moozXiv2XItlC12Gy6_LGq1t5mks1DQ`
	token = `t1.9euelZqKjJnIkJaWzpWVls_Hi8nIne3rnpWancvJj87JlImekoyYyozMk8rl8_d-WVdt-e9rbRNU_N3z9z4IVW3572ttE1T8zef1656VmpWemZzOl4uUmc-aj5vLysiL7_0.fnbS_pKMTWmhq9WRrcuVeutuMETM5NXjDQXrxSTCBDLoOxpcsp579Cj2ddjdV52SY0zhrToqXfaJbHt96ZDEAA`
)

const CHUNK_SIZE = 4000

// specification = stt_service_pb2.RecognitionSpec(
// 	language_code='ru-RU',
// 	profanity_filter=True,
// 	model='general',
// 	partial_results=True,
// 	audio_encoding='LINEAR16_PCM',
// 	sample_rate_hertz=8000
// )
// streaming_config = stt_service_pb2.RecognitionConfig(specification=specification, folder_id=folder_id)

// # Отправить сообщение с настройками распознавания.
// yield stt_service_pb2.StreamingRecognitionRequest(config=streaming_config)

// # Прочитать аудиофайл и отправить его содержимое порциями.
// with open(audio_file_name, 'rb') as f:
// 	data = f.read(CHUNK_SIZE)
// 	while data != b'':
// 			yield stt_service_pb2.StreamingRecognitionRequest(audio_content=data)
// 			data = f.read(CHUNK_SIZE)

func main() {
	logrus.Info("start")
	logrus.SetReportCaller(true)
	f, err := ioutil.ReadFile("/home/buzz/Downloads/t/test8K-30sec.pcm")
	if err != nil {
		logrus.Error(err)
		return
	}
	cred := credentials.NewTLS(&tls.Config{})
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(cred))
	conn, err := grpc.Dial("stt.api.cloud.yandex.net:443", opts...)
	if err != nil {
		logrus.Error("fail to dial: %v", err)
		return
	}
	defer conn.Close()

	cli := stt.NewSttServiceClient(conn)
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(map[string]string{"authorization": fmt.Sprintf("Bearer %s", token)}))
	specification := stt.RecognitionSpec{
		AudioEncoding:   stt.RecognitionSpec_LINEAR16_PCM,
		LanguageCode:    "ru-RU",
		ProfanityFilter: false,
		Model:           "general",
		PartialResults:  true,
		SampleRateHertz: 8000,
	}
	streaming_config := stt.RecognitionConfig{
		Specification: &specification,
		FolderId:      "b1g5k7e5erj5cbiudq71",
	}
	// stream, err := cli.StreamingRecognize(ctx)
	stream, err := cli.StreamingRecognize(ctx)

	// resp, err := cli.LongRunningRecognize(ctx, &stt.LongRunningRecognitionRequest{Config: &streaming_config, Audio: &stt.RecognitionAudio{AudioSource: &stt.RecognitionAudio_Uri{Uri: "https://storage.yandexcloud.net/duduhexchange/speech.pcm"}}})

	if err != nil {
		logrus.Error(err)
		return
	}
	// logrus.Infof("%+v", resp)
	// logrus.Info(resp.Result)
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				logrus.Info("Done")
				return
			}
			if err != nil {
				logrus.Error("Failed to receive a note : %v", err)
				return
			}
			logrus.Info(in.Chunks)
			for _, v := range in.Chunks {
				for _, vv := range v.Alternatives {
					logrus.Printf("Got message \n%s", vv.Text)
				}
			}
		}

	}()

	err = stream.Send(&stt.StreamingRecognitionRequest{
		StreamingRequest: &stt.StreamingRecognitionRequest_Config{Config: &streaming_config},
	})
	// Config: &streaming_config, Audio: &stt.RecognitionAudio{AudioSource: &stt.RecognitionAudio_Uri{Uri: "https://storage.yandexcloud.net/duduhexchange/test8K-30sec.mp3"}}})
	if err != nil {
		logrus.Error(err)
		return
	}
	r := bytes.NewBuffer(f)
	buf := make([]byte, 0, CHUNK_SIZE)
	// ticker := time.NewTicker(10 * time.Second)
	// err = stream.Send(&stt.StreamingRecognitionRequest{StreamingRequest: &stt.StreamingRecognitionRequest_AudioContent{AudioContent: f}})
	// if err != nil && err != io.EOF {
	// 	logrus.Error(err)
	// 	return
	// }
	for {
		n, err := r.Read(buf[:cap(buf)])
		buf = buf[:n]

		if n == 0 {
			if err == nil {
				break
			}
			if err == io.EOF {
				logrus.Info("sended")
				break
			}
			logrus.Error(err)
			return
		}
		// process buf
		// logrus.Info(len(buf))
		err = stream.Send(&stt.StreamingRecognitionRequest{StreamingRequest: &stt.StreamingRecognitionRequest_AudioContent{AudioContent: buf}})
		if err != nil && err != io.EOF {
			logrus.Error(err)
			return
		}
		// <-ticker.C
	}

	wg.Wait()

}

// #coding=utf8
// import argparse

// import grpc

// import yandex.cloud.ai.stt.v2.stt_service_pb2 as stt_service_pb2
// import yandex.cloud.ai.stt.v2.stt_service_pb2_grpc as stt_service_pb2_grpc

// CHUNK_SIZE = 4000

// def gen(folder_id, audio_file_name):
//     # Задать настройки распознавания.
//     specification = stt_service_pb2.RecognitionSpec(
//         language_code='ru-RU',
//         profanity_filter=True,
//         model='general',
//         partial_results=True,
//         audio_encoding='LINEAR16_PCM',
//         sample_rate_hertz=8000
//     )
//     streaming_config = stt_service_pb2.RecognitionConfig(specification=specification, folder_id=folder_id)

//     # Отправить сообщение с настройками распознавания.
//     yield stt_service_pb2.StreamingRecognitionRequest(config=streaming_config)

//     # Прочитать аудиофайл и отправить его содержимое порциями.
//     with open(audio_file_name, 'rb') as f:
//         data = f.read(CHUNK_SIZE)
//         while data != b'':
//             yield stt_service_pb2.StreamingRecognitionRequest(audio_content=data)
//             data = f.read(CHUNK_SIZE)

// def run(folder_id, iam_token, audio_file_name):
//     # Установить соединение с сервером.
//     cred = grpc.ssl_channel_credentials()
//     channel = grpc.secure_channel('stt.api.cloud.yandex.net:443', cred)
//     stub = stt_service_pb2_grpc.SttServiceStub(channel)

//     # Отправить данные для распознавания.
//     it = stub.StreamingRecognize(gen(folder_id, audio_file_name), metadata=(('authorization', 'Bearer %s' % iam_token),))

//     # Обработать ответы сервера и вывести результат в консоль.
//     try:
//         for r in it:
//             try:
//                 print('Start chunk: ')
//                 for alternative in r.chunks[0].alternatives:
//                     print('alternative: ', alternative.text)
//                 print('Is final: ', r.chunks[0].final)
//                 print('')
//             except LookupError:
//                 print('Not available chunks')
//     except grpc._channel._Rendezvous as err:
//         print('Error code %s, message: %s' % (err._state.code, err._state.details))

// if __name__ == '__main__':
//     parser = argparse.ArgumentParser()
//     parser.add_argument('--token', required=True, help='IAM token')
//     parser.add_argument('--folder_id', required=True, help='folder ID')
//     parser.add_argument('--path', required=True, help='audio file path')
//     args = parser.parse_args()

//     run(args.folder_id, args.token, args.path)
