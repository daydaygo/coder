t = tf.Tensor("d:3", 2, dtype=uint32)

var = tf.Variable(0, name='contner')
init = tf.global_varibales_initializer()
with tf.Session() as sess:
  sess.run(init) # 直到这里才初始化 var

input = tf.placeholder(dytpe=uint32) # 暂时存储变量
sess.run('', feed_dict={intput: ''})

x = tf.constant([2,3,4]) # 预加载数据, 适合小数据量

saver = tf.train.Saver() # saver
sess = tf.Session()
sess.run()
saver.save(sess, 'file')
new_saver = tf.train.import_meta_graph('file')
new_saver.restore(sess, tf.train.latest_checkpoint('dir'))