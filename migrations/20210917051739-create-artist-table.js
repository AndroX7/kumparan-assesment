'use strict';

module.exports = {
  up: async (queryInterface, Sequelize) => {
    return queryInterface.createTable('articles', {
        id: {
          type: Sequelize.BIGINT,
          allowNull: false,
          primaryKey: true,
          autoIncrement: true
      },
      author:{
          type: Sequelize.STRING,
          allowNull: true,
      },
      author_id:{
        type: Sequelize.BIGINT,
        allowNull: true,
    },
      title:{
        type: Sequelize.STRING,
        allowNull: true,
      },
      body: {
        type: Sequelize.TEXT,
        allowNull: true
      },
      created: {
        type: Sequelize.BIGINT,
        allowNull: true
      }, 
      created_at: {
        type: Sequelize.DATE,
        allowNull: false
      },
      updated_at: {
          type: Sequelize.DATE,
          allowNull: false
      },
      deleted_at: {
          type: Sequelize.DATE,
          allowNull: true
      }
    })
  },

  down: async (queryInterface, Sequelize) => {
    return queryInterface.dropTable('artists');
  }
};
