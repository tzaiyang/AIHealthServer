import time
from kafka import KafkaConsumer, TopicPartition


class GetOffsetWithTimestamp:
    def __init__(self, broker_list, topic):
        self.topic = topic
        self.consumer = KafkaConsumer(topic,bootstrap_servers=broker_list)

    def get_offset_time_window(self, begin_time, end_time):
        partitions_structs = []

        for partition_id in self.consumer.partitions_for_topic(self.topic):
            partitions_structs.append(TopicPartition(self.topic, partition_id))

        begin_search = {}
        for partition in partitions_structs:
            begin_search[partition] = begin_time if isinstance(begin_time, int) else self.str_to_timestamp(begin_time)
        begin_offset = self.consumer.offsets_for_times(begin_search)

        end_search = {}
        for partition in partitions_structs:
            end_search[partition] = end_time if isinstance(end_time, int) else self.str_to_timestamp(end_time)
        end_offset = self.consumer.offsets_for_times(end_search)

        for topic_partition, offset_and_timestamp in begin_offset.items():
            b_offset = 'null' if offset_and_timestamp is None else offset_and_timestamp[0]
            e_offset = 'null' if end_offset[topic_partition] is None else end_offset[topic_partition][0]
            print('Between {0} and {1}, {2} offset range = [{3}, {4}]'.format(begin_time, end_time, topic_partition,b_offset, e_offset))
            self.consumer.seek(topic_partition,b_offset)

        return self.consumer,e_offset


    def str_to_timestamp(self, str_time, format_type="%Y-%m-%d %H:%M:%S"):
        time_array = time.strptime(str_time, format_type)
        return int(time.mktime(time_array)) * 1000


if __name__ == '__main__':
    broker_list = 'localhost:9092'
    topic = 'access_log'

    action = GetOffsetWithTimestamp(broker_list, topic)
    consumer,e_offset=action.get_offset_time_window('2021-05-30 11:00:00', '2021-05-30 12:30:00')

    print(consumer.poll(50).values())

    # for message in action.consumer:   
    #   # logger.info("%s value=%s" %(message.topic,message.value))
    #   print ("%s %s value=%s" % (time.strftime('%Y-%m-%d %H:%M:%S', time.localtime(time.time())),message.topic, message.value))