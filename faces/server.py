
from concurrent import futures
import time
import grpc
import recognizer_pb2_grpc as grpc_proto
import recognizer_pb2 as proto
import os
import recognizer


class Servicer(grpc_proto.FaceRecognizerServicer):
    def RecognizeFaces(self, request, context):

        faces = recognizer.recognize(request.url)
        proto_faces = []
        for face in faces:
            area = face['area']
            encoding = face['encoding']
            proto_face = proto.Face(x=area['x'], y=area['y'], width=area['width'],
                                    height=area['height'], encoding=encoding)
            proto_faces.append(proto_face)

        return proto.RegognizeResponse(faces=proto_faces)


def serve():

    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    grpc_proto.add_FaceRecognizerServicer_to_server(
        Servicer(),
        server
    )

    server.add_insecure_port('[::]:' + os.environ['GRPC_PORT'])

    server.start()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)


if __name__ == '__main__':
    serve()
